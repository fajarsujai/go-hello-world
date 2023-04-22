pipeline {
    agent any

    stages{
        stage("Docker Build"){
            steps{
                script{
                    sh "docker build -t fajarsujai/go-helloworld:${GIT_COMMIT} ."
                }
            }
        }


        stage("Docker Push"){
            steps{
                script{
                    sh "docker push fajarsujai/go-helloworld:${GIT_COMMIT}"
                }
            }
        }


        stage("Finish"){
            steps{
                script{
                    sh "echo('berhsil push image docker')"
                }
            }
        }
    }
}