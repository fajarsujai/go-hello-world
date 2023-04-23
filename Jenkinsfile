pipeline {
    agent any

    stages{
        // stage("Docker Build"){
        //     steps{
        //         script{
        //             sh "docker build -t fajarsujai/go-helloworld:${GIT_COMMIT} --build-arg BRANCH=develop --build-arg PORT=8001 ."
        //         }
        //     }
        // }


        // stage("Docker Push"){
        //     steps{
        //         script{
        //             sh "docker push fajarsujai/go-helloworld:${GIT_COMMIT}"
        //         }
        //     }
        // }


        // stage("Finish"){
        //     steps{
        //         script{
        //             echo("berhsil push image docker")
        //         }
        //     }
        // }

        stage("Ansible version"){
            steps{
                script{
                    sh "exec -it 3eb20105ab7c ansible --version"
                }
            }
        }

        stage("Docker version"){
            steps{
                script{
                    sh "docker version"
                }
            }
        }
    }
}