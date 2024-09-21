package com.phuc.productservice.configuations;

import com.corundumstudio.socketio.SocketIOServer;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class SocketIOConfiguration {

    @Bean
    public SocketIOServer socketIOServer() {
        com.corundumstudio.socketio.Configuration socketIOConfig = new com.corundumstudio.socketio.Configuration();
        socketIOConfig.setHostname("localhost");
        socketIOConfig.setPort(7000);

        return new SocketIOServer(socketIOConfig);
    }
}
