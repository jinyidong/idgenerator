FROM harbor.suiyi.com.cn/monitor/alpine:latest
ADD IdGenerator /
CMD ["/IdGenerator"]