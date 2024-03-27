/*
Copyright 2018 The kube-fledged authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package images

import (
	"os"
	"testing"

	fledgedv1alpha2 "github.com/senthilrch/kube-fledged/pkg/apis/kubefledged/v1alpha2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var imageCache = fledgedv1alpha2.ImageCache{
	ObjectMeta: metav1.ObjectMeta{
		Name:      "foo",
		Namespace: "kube-fledged",
	},
	Spec: fledgedv1alpha2.ImageCacheSpec{
		CacheSpec: []fledgedv1alpha2.CacheSpecImages{
			{
				Images: []string{"foo"},
			},
		},
	},
}

func TestTolerationsNotPresentByDefault(t *testing.T) {
	jobspec, _ := newImagePullJob(&imageCache, "foo", &node, "IfNotPresetn", "busyboxImage", "fooSA", "default")
	if len(jobspec.Spec.Template.Spec.Tolerations) > 0 {
		t.Errorf("Test Failed, Expected No Tolerations, Found: %v", jobspec.Spec.Template.Spec.Tolerations)
	}
}

func TestTolerationsPresentOnEnvSet(t *testing.T) {
	os.Setenv("SET_JOB_TOLERATIONS", "yes")
	jobspec, _ := newImagePullJob(&imageCache, "foo", &node, "IfNotPresetn", "busyboxImage", "fooSA", "default")
	if len(jobspec.Spec.Template.Spec.Tolerations) == 0 {
		t.Errorf("Test Failed, Expected Tolerations, None Found: %v", jobspec.Spec.Template.Spec.Tolerations)
	}
}
