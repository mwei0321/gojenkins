pipeline {
	agent any
	
	environment {
		key = 'value'
	}
	stages {
		stage('拉取仓库代码') {
            steps {
                echo '摘取代码成功'
            }
		}
        stage('编译文件') {
            steps {
                echo '编译文件成功'
            }
		}
        stage('打包docker镜像') {
            steps {
                echo '打包docker镜像成功'
            }
		}
	}
}