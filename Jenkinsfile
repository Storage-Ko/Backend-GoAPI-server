node {
    git poll: true, url:'https://github.com/Ko-GyeongTae/Backend-GoAPI-server.git'
        { 
            stage('Pull') {
                git 'https://github.com/Ko-GyeongTae/Backend-GoAPI-server.git' 
            } 
            stage('Unit Test') { 
            
            } 
            stage('Build') { 
                sh(script: 'go build -o main .')  
            } 
            stage('Deploy') { 
                sh(script: 'docker-compose up -d production')    
            } 
        } 
    
}
