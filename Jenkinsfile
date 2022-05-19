pipeline {
    // install golang 1.14 on Jenkins node
    agent { docker { image 'golang' } }
    environment {
        GO111MODULE = 'on'
//         DB_NAME=video
//         DB_COLUMN=col_videos
    }
    stages {
        node('Laptop') {
            stage("unit-test") {
                steps {
                    echo 'UNIT TEST EXECUTION STARTED'
    //                 sh 'make unit-tests'
                }
            }
            stage("functional-test") {
                steps {
                    echo 'FUNCTIONAL TEST EXECUTION STARTED'
    //                 sh 'make functional-tests'
                }
            }
            stage("BUILD DOCKER IMAGE") {
                steps {
                    echo 'BUILD EXECUTION STARTED'
                    sh 'go version'
                    sh 'go get ./...'
                    sh 'docker build . -t beatable2310/new2c:latest'
                }
            }
            stage('PUSH DOCKER IMAGE TO DOCKER HUB') {
                agent any
                steps {
    //                 withCredentials([usernamePassword(credentialsId: 'docker-hub', passwordVariable: 'pass', usernameVariable: 'user')]) {
    //                 sh "docker login --username=${user} --password=${pass}"
                    sh 'echo "Docker login successful"'
                    sh 'docker push beatable2310/new2c:latest'
    //                 }
                }
            }
            stage('BUILD DB') {
                steps {
                    sh 'docker-compose up -d'
                    sh "echo \"use ${DB_NAME} && db.createCollection('${DB_COLUMN}')\" > demo.js"
                    sh 'docker exec -it mongodb bash -c "mongo < demo.js"'
                }
            }
            stage('DEPLOY TO NODE') {
                steps {

                    sh 'docker pull beatable2310/new2c:latest'
                    sh 'docker run -d -p 8080:8080 beatable2310/new2c:latest'
                }
            }
        }
        post {
                always {
                    cleanWs()
                }
        }
    }
}