// Copyright 2019 The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package channel

import (
	"testing"
	"time"

	"github.com/onsi/gomega"
	"golang.org/x/net/context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	chv1 "github.com/open-cluster-management/multicloud-operators-channel/pkg/apis/apps/v1"
)

var c client.Client

const (
	timeout           = time.Second * 5
	targetNamespace   = "default"
	tragetChannelName = "foo"
	targetChannelType = chv1.ChannelType("namespace")
)

var expectedRequest = reconcile.Request{NamespacedName: types.NamespacedName{Name: tragetChannelName, Namespace: targetNamespace}}

func TestChannelControllerReconcile(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	channelInstance := &chv1.Channel{
		ObjectMeta: metav1.ObjectMeta{Name: tragetChannelName, Namespace: targetNamespace},
		Spec: chv1.ChannelSpec{
			Type:     targetChannelType,
			Pathname: targetNamespace,
		},
	}

	// Setup the Manager and Controller.  Wrap the Controller Reconcile function so it writes each request to a
	// channel when it is finished.
	mgr, err := manager.New(cfg, manager.Options{MetricsBindAddress: "0"})
	g.Expect(err).NotTo(gomega.HaveOccurred())

	c = mgr.GetClient()

	//create events handler on hub cluster. All the deployable events will be written to the root deployable on hub cluster.
	hubClientSet, _ := kubernetes.NewForConfig(cfg)

	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(klog.Infof)
	eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{Interface: hubClientSet.CoreV1().Events("")})
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, corev1.EventSource{Component: "channel"})

	recFn, requests := SetupTestReconcile(newReconciler(mgr, recorder))
	g.Expect(add(mgr, recFn)).NotTo(gomega.HaveOccurred())

	stopMgr, mgrStopped := StartTestManager(mgr, g)

	defer func() {
		close(stopMgr)
		mgrStopped.Wait()
	}()

	// Create the Channel object and expect the Reconcile
	err = c.Create(context.TODO(), channelInstance)

	g.Expect(err).NotTo(gomega.HaveOccurred())

	defer c.Delete(context.TODO(), channelInstance)

	time.Sleep(time.Second * 1)
	g.Eventually(requests, timeout).Should(gomega.Receive(gomega.Equal(expectedRequest)))
}

// test if referred secret and configmap were annotated correctly or not
func TestChannelAnnotateReferredSecertAndConfigMap(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	refSrtName := "ch-srt"
	refSrt := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      refSrtName,
			Namespace: targetNamespace,
		},
	}

	refCmName := "ch-cm"
	refCm := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      refCmName,
			Namespace: targetNamespace,
		},
	}

	chKey := types.NamespacedName{Name: tragetChannelName, Namespace: targetNamespace}
	chn := &chv1.Channel{
		ObjectMeta: metav1.ObjectMeta{Name: tragetChannelName, Namespace: targetNamespace},
		Spec: chv1.ChannelSpec{
			Type:         targetChannelType,
			Pathname:     targetNamespace,
			SecretRef:    &corev1.ObjectReference{Name: refSrtName, Namespace: targetNamespace, Kind: "secret"},
			ConfigMapRef: &corev1.ObjectReference{Name: refCmName, Namespace: targetNamespace, Kind: "ConfigMap"},
		},
	}

	// Setup the Manager and Controller.  Wrap the Controller Reconcile function so it writes each request to a
	// channel when it is finished.
	mgr, err := manager.New(cfg, manager.Options{MetricsBindAddress: "0"})
	g.Expect(err).NotTo(gomega.HaveOccurred())

	c = mgr.GetClient()

	tRecorder := record.NewBroadcaster().NewRecorder(mgr.GetScheme(), corev1.EventSource{Component: "channel"})
	// g.Expect(Add(mgr, tRecorder, nil, nil, nil)).NotTo(gomega.HaveOccurred())

	stopMgr, mgrStopped := StartTestManager(mgr, g)

	defer func() {
		close(stopMgr)
		mgrStopped.Wait()
	}()

	rec := newReconciler(mgr, tRecorder)

	defer c.Delete(context.TODO(), refSrt)
	g.Expect(c.Create(context.TODO(), refSrt)).NotTo(gomega.HaveOccurred())

	defer c.Delete(context.TODO(), refCm)
	g.Expect(c.Create(context.TODO(), refCm)).NotTo(gomega.HaveOccurred())

	defer c.Delete(context.TODO(), chn)
	g.Expect(c.Create(context.TODO(), chn)).NotTo(gomega.HaveOccurred())

	rq := reconcile.Request{NamespacedName: chKey}

	_, err = rec.Reconcile(rq)
	g.Expect(err).NotTo(gomega.HaveOccurred())

	updatedSrt := &corev1.Secret{}
	g.Expect(c.Get(
		context.TODO(),
		types.NamespacedName{Name: refSrtName, Namespace: targetNamespace},
		updatedSrt)).NotTo(gomega.HaveOccurred())

	assertReferredObjAnno(t, updatedSrt, chKey.String())

	updatedCm := &corev1.ConfigMap{}
	g.Expect(c.Get(
		context.TODO(),
		types.NamespacedName{Name: refCmName, Namespace: targetNamespace},
		updatedCm)).NotTo(gomega.HaveOccurred())

	assertReferredObjAnno(t, updatedCm, chKey.String())
}

func assertReferredObjAnno(t *testing.T, obj metav1.Object, chKeyStr string) {
	anno := obj.GetAnnotations()
	if len(anno) == 0 {
		t.Errorf("anno is empty")
		return
	}

	if anno[chv1.ServingChannel] == "" {
		t.Errorf("target annotation is missing")
		return
	}

	if anno[chv1.ServingChannel] != chKeyStr {
		t.Errorf("referred annotation, wanted %v got %v\n", chKeyStr, anno[chv1.ServingChannel])
		return
	}
}

// 1. test role and rolebinding set up
// 2. test the annotation deletion
