/*
Copyright 2019 Replicated, Inc..

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
// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"time"

	v1beta1 "github.com/replicatedhq/troubleshoot/pkg/apis/troubleshoot/v1beta1"
	scheme "github.com/replicatedhq/troubleshoot/pkg/client/troubleshootclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AnalyzerJobsGetter has a method to return a AnalyzerJobInterface.
// A group's client should implement this interface.
type AnalyzerJobsGetter interface {
	AnalyzerJobs(namespace string) AnalyzerJobInterface
}

// AnalyzerJobInterface has methods to work with AnalyzerJob resources.
type AnalyzerJobInterface interface {
	Create(*v1beta1.AnalyzerJob) (*v1beta1.AnalyzerJob, error)
	Update(*v1beta1.AnalyzerJob) (*v1beta1.AnalyzerJob, error)
	UpdateStatus(*v1beta1.AnalyzerJob) (*v1beta1.AnalyzerJob, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.AnalyzerJob, error)
	List(opts v1.ListOptions) (*v1beta1.AnalyzerJobList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.AnalyzerJob, err error)
	AnalyzerJobExpansion
}

// analyzerJobs implements AnalyzerJobInterface
type analyzerJobs struct {
	client rest.Interface
	ns     string
}

// newAnalyzerJobs returns a AnalyzerJobs
func newAnalyzerJobs(c *TroubleshootV1beta1Client, namespace string) *analyzerJobs {
	return &analyzerJobs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the analyzerJob, and returns the corresponding analyzerJob object, and an error if there is any.
func (c *analyzerJobs) Get(name string, options v1.GetOptions) (result *v1beta1.AnalyzerJob, err error) {
	result = &v1beta1.AnalyzerJob{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("analyzerjobs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AnalyzerJobs that match those selectors.
func (c *analyzerJobs) List(opts v1.ListOptions) (result *v1beta1.AnalyzerJobList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.AnalyzerJobList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("analyzerjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested analyzerJobs.
func (c *analyzerJobs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("analyzerjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a analyzerJob and creates it.  Returns the server's representation of the analyzerJob, and an error, if there is any.
func (c *analyzerJobs) Create(analyzerJob *v1beta1.AnalyzerJob) (result *v1beta1.AnalyzerJob, err error) {
	result = &v1beta1.AnalyzerJob{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("analyzerjobs").
		Body(analyzerJob).
		Do().
		Into(result)
	return
}

// Update takes the representation of a analyzerJob and updates it. Returns the server's representation of the analyzerJob, and an error, if there is any.
func (c *analyzerJobs) Update(analyzerJob *v1beta1.AnalyzerJob) (result *v1beta1.AnalyzerJob, err error) {
	result = &v1beta1.AnalyzerJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("analyzerjobs").
		Name(analyzerJob.Name).
		Body(analyzerJob).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *analyzerJobs) UpdateStatus(analyzerJob *v1beta1.AnalyzerJob) (result *v1beta1.AnalyzerJob, err error) {
	result = &v1beta1.AnalyzerJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("analyzerjobs").
		Name(analyzerJob.Name).
		SubResource("status").
		Body(analyzerJob).
		Do().
		Into(result)
	return
}

// Delete takes name of the analyzerJob and deletes it. Returns an error if one occurs.
func (c *analyzerJobs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("analyzerjobs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *analyzerJobs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("analyzerjobs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched analyzerJob.
func (c *analyzerJobs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.AnalyzerJob, err error) {
	result = &v1beta1.AnalyzerJob{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("analyzerjobs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
