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

=> [internal] load build definition from Dockerfile

=> => transferring dockerfile: 31B     

=> [internal] load .dockerignore   

=> => transferring context: 2B  

=> [internal] load metadata for docker.io/library/golang:1.15-alpine3.14 

=> [1/5] FROM docker.io/library/golang:1.15-alpine3.14@sha256:5b8a2dd7cbf2d0f59906d8cc4bff338d73faa481b492187863ab72839037b2cc 

=> [internal] load build context  

=> => transferring context: 8.64kB    

=> CACHED [2/5] RUN mkdir /app      

=> [3/5] ADD . /app            

=> [4/5] WORKDIR /app      

=> [5/5] RUN go build -o main .   

=> exporting to image    

=> => exporting layers    

=> => writing image sha256:edc6f3da84154f03a5953165946e3358cc0f70288847c3f051b0754166ea7909 

=> => naming to docker.io/library/mygolangdemo:1.15    


# K8's file will be used apply the changes to deploy the mygolangdemo image with Tag as 1.15 to POD

kubectl apply -f C:\Workspace\golangdemo\K8S.yaml





