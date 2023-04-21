pipeline {
    agent any

    stages{
        stage("Docker Version"){
            steps{
                script{
                    sh "docker version"
                }
            }
        }


        stage("Kubectl Version"){
            steps{
                script{
                    sh "kubectl version"
                }
            }
        }
    }
}