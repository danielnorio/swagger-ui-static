FROM golang as compiler 
RUN CGO_ENABLED=0 go get -a -ldflags '-s' \ 
  github.com/danielnorio/swagger-ui-static  

FROM scratch
COPY --from=compiler /go/bin/swagger-ui-static .
ADD /static ./static
EXPOSE 3000
CMD ["./swagger-ui-static"]