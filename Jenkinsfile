pipeline {
  agent none
  stages {
    stage('ShellScript') {
      parallel {
        stage('ShellScript') {
          steps {
            sh '''#!/bin/bash

echo "Ol�"'''
          }
        }

        stage('ShellScript2') {
          steps {
            sh '''#!/bin/bash

echo "Ol� 2"'''
          }
        }

      }
    }

  }
}