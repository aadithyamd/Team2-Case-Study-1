
// pipeline {
//     agent { docker { image 'golang' } }
//     stages {
//         stage('build') {
//             steps {
//                 sh 'ls -la'
//                 sh 'pwd'
//                 sh 'go version'
//                 sh 'export XDG_CACHE_HOME=/tmp/.cache'
//                 sh 'echo \$XDG_CACHE_HOME'
//                 sh 'sudo go mod download'
//                 sh 'sudo make api'
//                 sh 'sudo make server'
//             }
//         }
//     }
// }


pipeline {

    agent any

    tools {

        go 'go1.14'

    }

    stages {

        stage('Build and deploy docker image') {

            steps {

                echo 'Building new gin image'
                sh 'docker build -t gin-t2 -f Dockerfile.gin .'

                echo 'Building new grpc image'
                sh 'docker build -t grpc-t2 -f Dockerfile.grpc .'

                echo 'Stopping existing gin container'
                sh 'docker stop container-go-t2'

                echo 'Stopping existing gin container'
                sh 'docker stop container-grpc-t2'

                sh 'docker run --rm -dp 5051:5051 --net=mynet --name=container-grpc-t2 grpc-t2'

                sh 'docker run --rm -dp 9001:9001 --net=mynet --name=container-gin-t2 gin-t2'



            }

        }