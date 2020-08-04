// pipeline {
//     agent any
//     tools {
//        go 'go1.14'
//     }
//     stages {
//         stage('Pre Test') {
//             steps {
//                 echo 'Installing dependencies'
//                 sh 'ls -la'
//                 sh 'pwd'
//                 go version
//             }
//         }
//
//         stage('Execute') {
//             steps {
//                 echo 'Compiling and building'
//                 sh 'make api'
//                 sh 'make server'
//             }
//         }
//     }
// }

pipeline {
    agent { docker { image 'golang' } }
    stages {
        stage('build') {
            steps {
                sh 'go version'
                sh 'go mod download'
                sh 'make api'
                sh 'make server'
            }
        }
    }
}