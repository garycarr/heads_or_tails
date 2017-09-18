// https://jenkins.io/doc/book/pipeline/jenkinsfile/

pipeline {
    agent any

    parameters {
        string(name: 'env_colour', defaultValue: 'black', description: 'An example of how we could use parametized builds with envs if we wanted')
    }
    environment {
        APP_NAME = "heads_or_tails"
        BUILD_DIR = "_build/src/github.com/garycarr"
        COLOUR = ${params.env_colour} // Set any correctly spelled environmental we want to use
        DOCKER_IMAGE_ID = "turnitin/heads_or_tails:${env.BUILD_NUMBER}"
        DOCKER_CONTAINER_ID = "heads_or_tails_${env.BUILD_NUMBER}"
    }

    stages {
        stage('Build go repo and test') {
            steps {
                sh './build_scripts/build_go_repo_and_test.sh'
            }
        }
        stage('Build docker image and test') {
            steps {
                sh './build_scripts/build_docker_image_and_test.sh'
            }
        }
        stage('Deploy') {
            steps {
                script {
                    // Need to decide what happens before deploy.sh as we need the kube deploy repo
                    // Git submodule it? Clone it?
                    println("BRANCH NAME IS ${env.BRANCH_NAME}")
                    // Will this even work on PR's?
                    if (env.BRANCH_NAME == 'master') {
                        sh './build_scripts/deploy.sh prod'
                    } else {
                        // What about these builds in gcloud
                        sh './build_scripts/deploy.sh dev'
                        sh './build_scripts/deploy.sh qa'
                    }
                }
            }
        }
    }
    post {
        // Like a finally
        always {
            sh './build_scripts/clean_up.sh'
        }
        success {
            sh '''
                echo "This always runs on success"
                echo "And is an example of multiline"
            '''
        }
        failure {
            // This runs on failure
            echo "We have failed"
        }
    }
}
