pipeline {
  agent any
  stages {
    stage('Build & Push') {
      steps {
        sh '''docker build -t anilkuscu95/demo:$BUILD_NUMBER .
docker tag anilkuscu95/demo:$BUILD_NUMBER anilkuscu95/demo:latest
docker login --username=$DOCKER_USERNAME --password=$DOCKER_PASSWORD
docker push anilkuscu95/demo:$BUILD_NUMBER
docker push anilkuscu95/demo:latest'''
      }
    }

    stage('Deploy to Test Environment') {
      steps {
        sh '''sed -i "s|image: .*$|image: anilkuscu95/demo:${BUILD_NUMBER}|" $YAML_FILE
kubectl apply -f $YAML_FILE'''
      }
    }

    stage('Testing') {
      steps {
        sh '''sed -i "s|  name: demo-test-job |  name: demo-test-job-${BUILD_NUMBER}|" $TEST_YAML
kubectl apply -f $TEST_YAML
while true; do [ "$(kubectl get po | grep "demo-test-job-${BUILD_NUMBER}" | awk \'{print $3}\')" = "Completed" ] && break; done
TEST_RESULT=$(kubectl logs $(kubectl get po | grep "demo-test-job-${BUILD_NUMBER}" | awk \'{print $1}\'))
if [[ $TEST_RESULT == "exit_code:0" ]];then
exit 0
else
exit 1
fi'''
      }
    }

stage('Approval') {
agent none
options {
timeout(time: 20, unit: "MINUTES")
}      
steps {
script {
env.DEPLOY = input message: 'Click Yes if You Want to Contunie Deployment',
parameters: [choice(name: 'Do you want to deploy it to production ?', choices: 'no\nyes', description: 'Choose "yes" if you want to deploy this build to production')]
}
}
}

    stage('Deploy to Production') {
      steps {
        sh '''sed -i "s|image: .*$|image: anilkuscu95/demo:${BUILD_NUMBER}|" $YAML_FILE
sed -i "s|namespace: .*$|namespace: production |" $YAML_FILE
kubectl apply -f $YAML_FILE'''
      }
    }

  }
  environment {
    DOCKER_USERNAME = 'anilkuscu95'
    DOCKER_PASSWORD = '<MyPass>'
    YAML_FILE = 'demo.yaml'
    TEST_YAML = 'test.yaml'
  }
}