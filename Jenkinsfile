pipeline {
	agent any
	
	environment {
		key = 'value'
	}
	stages {
		stage('拉取仓库代码') {
            steps {
                git 'https://github.com/mwei0321/gojenkins.git'
            }
		}
        stage('编译文件') {
            steps {
                sh 'go build -o ./main ./main.go'
            }
		}
        stage('打包docker镜像') {
            steps {
                echo '打包docker镜像成功'
            }
		}
	}
}