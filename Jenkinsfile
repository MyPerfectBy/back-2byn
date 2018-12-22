pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                 sh "go get -d ./..."
            }
        }
//        stage('Delivery') {
//            steps {
//                 sh "sudo -u root /usr/bin/php7.1 bin/console doctrine:migrations:migrate"
//                 sh "sudo -u root rm -rf /home/Assembly/Prototype/${branch}/Back/Code/*[^web]"
//                 sh "cp -r * /home/Assembly/Prototype/${branch}/Back/Code/"
//            }
//        }
    }
}
