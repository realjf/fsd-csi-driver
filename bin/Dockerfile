FROM alpine

LABEL maintainers="realjf"
LABEL description="FSD CSI Driver"

RUN apk add util-linux e2fsprogs
COPY fsd-csi-driver /fsd-csi-driver 

ENTRYPOINT ["/fsd-csi-driver"]
