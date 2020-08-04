pipeline {
    agent any
//     tools {
//     //    go 'go1.14'
//     }
    stages {        
        stage('Pre Test') {
            steps {
                echo 'Installing dependencies'
                sh 'ls -la'
                sh 'pwd'
                //sh 'go version'
            }
        }
        
        stage('Execute') {
            steps {
                echo 'Compiling and building'

                node {
                    // Install the desired Go version
                    def root = tool name: 'Go 1.8', type: 'go'

                    // Export environment variables pointing to the directory where Go was installed
                    withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
                        sh 'go version'
                        sh 'make api'
                        sh 'make server'

                    }
                }
            }
        }
}
}