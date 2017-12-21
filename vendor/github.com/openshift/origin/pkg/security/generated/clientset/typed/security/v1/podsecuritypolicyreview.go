package v1

import (
	v1 "github.com/openshift/origin/pkg/security/apis/security/v1"
	rest "k8s.io/client-go/rest"
)

// PodSecurityPolicyReviewsGetter has a method to return a PodSecurityPolicyReviewInterface.
// A group's client should implement this interface.
type PodSecurityPolicyReviewsGetter interface {
	PodSecurityPolicyReviews(namespace string) PodSecurityPolicyReviewInterface
}

// PodSecurityPolicyReviewInterface has methods to work with PodSecurityPolicyReview resources.
type PodSecurityPolicyReviewInterface interface {
	Create(*v1.PodSecurityPolicyReview) (*v1.PodSecurityPolicyReview, error)
	PodSecurityPolicyReviewExpansion
}

// podSecurityPolicyReviews implements PodSecurityPolicyReviewInterface
type podSecurityPolicyReviews struct {
	client rest.Interface
	ns     string
}

// newPodSecurityPolicyReviews returns a PodSecurityPolicyReviews
func newPodSecurityPolicyReviews(c *SecurityV1Client, namespace string) *podSecurityPolicyReviews {
	return &podSecurityPolicyReviews{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Create takes the representation of a podSecurityPolicyReview and creates it.  Returns the server's representation of the podSecurityPolicyReview, and an error, if there is any.
func (c *podSecurityPolicyReviews) Create(podSecurityPolicyReview *v1.PodSecurityPolicyReview) (result *v1.PodSecurityPolicyReview, err error) {
	result = &v1.PodSecurityPolicyReview{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("podsecuritypolicyreviews").
		Body(podSecurityPolicyReview).
		Do().
		Into(result)
	return
}
