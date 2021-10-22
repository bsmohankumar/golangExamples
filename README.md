# golangdemo

# Execution for main program as below 
PS C:\Workspace\golangdemo> go run .\main.go
Hello Welcome to Golng World!

# Execution for SRC folder as below 

PS C:\Workspace\golangdemo> go run .\src\ArraysExamples.go
Grades [97 98 99]
Grades1 [97 98 99 100]
students [  ]
students [Bhuvan BM Mohan Kumar BS Beena CB]
Number of students 3
Matrix [[1 2 3] [4 5 6] [7 8 9]]

# Docker steps : this will create mygolangdemo image with Tag as 1.15

PS C:\Workspace\golangdemo> docker build -t mygolangdemo:1.15 -f .\Dockerfile .
[+] Building 6.8s (10/10) FINISHED
=> [internal] load build definition from Dockerfile                                                                                                                               0.0s
=> => transferring dockerfile: 31B                                                                                                                                                0.0s
=> [internal] load .dockerignore                                                                                                                                                  0.0s
=> => transferring context: 2B                                                                                                                                                    0.0s
=> [internal] load metadata for docker.io/library/golang:1.15-alpine3.14                                                                                                          2.7s
=> [1/5] FROM docker.io/library/golang:1.15-alpine3.14@sha256:5b8a2dd7cbf2d0f59906d8cc4bff338d73faa481b492187863ab72839037b2cc                                                    0.0s
=> [internal] load build context                                                                                                                                                  0.1s
=> => transferring context: 8.64kB                                                                                                                                                0.0s
=> CACHED [2/5] RUN mkdir /app                                                                                                                                                    0.0s
=> [3/5] ADD . /app                                                                                                                                                               0.3s
=> [4/5] WORKDIR /app                                                                                                                                                             0.1s
=> [5/5] RUN go build -o main .                                                                                                                                                   3.3s
=> exporting to image                                                                                                                                                             0.2s
=> => exporting layers                                                                                                                                                            0.1s
=> => writing image sha256:edc6f3da84154f03a5953165946e3358cc0f70288847c3f051b0754166ea7909                                                                                       0.0s
=> => naming to docker.io/library/mygolangdemo:1.15      

#K8's file will be used apply the changes to deploy the mygolangdemo image with Tag as 1.15 to POD

kubectl apply -f C:\Workspace\golangdemo\K8S.yaml





