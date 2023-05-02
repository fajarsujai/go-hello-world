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

        stage("Run ansible playbook"){
            steps{
                script{
                   ansiblePlaybook(credentialsId: '-----BEGIN OPENSSH PRIVATE KEY----- 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-----END OPENSSH PRIVATE KEY-----', inventory: 'inventory/inventory.yaml', playbook: 'playbook.yml')
                }
            }
        }

        // stage("Docker version"){
        //     steps{
        //         script{
        //             sh "docker version"
        //         }
        //     }
        // }
    }
}