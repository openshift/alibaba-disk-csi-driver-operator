FROM registry.ci.openshift.org/ocp/builder:rhel-9-golang-1.21-openshift-4.16 AS builder
WORKDIR /go/src/github.com/openshift/alibaba-disk-csi-driver-operator
COPY . .
RUN make

FROM registry.ci.openshift.org/ocp/4.16:base-rhel9
COPY --from=builder /go/src/github.com/openshift/alibaba-disk-csi-driver-operator/alibaba-disk-csi-driver-operator /usr/bin/
ENTRYPOINT ["/usr/bin/alibaba-disk-csi-driver-operator"]
LABEL io.k8s.display-name="OpenShift Alibaba Disk CSI Driver Operator" \
	io.k8s.description="The Alibala Disk CSI Driver Operator installs and maintains the Alibaba Disk CSI Driver on a cluster."
