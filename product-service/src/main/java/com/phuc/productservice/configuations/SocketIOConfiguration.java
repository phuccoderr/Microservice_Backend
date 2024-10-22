package com.phuc.productservice.configuations;

import com.corundumstudio.socketio.SocketIOClient;
import com.corundumstudio.socketio.SocketIOServer;
import com.corundumstudio.socketio.Transport;
import com.corundumstudio.socketio.listener.ConnectListener;
import jakarta.annotation.PreDestroy;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class SocketIOConfiguration {

    private SocketIOServer server;

    @Bean
    public SocketIOServer SocketIOConfiguration() {
        com.corundumstudio.socketio.Configuration socketIOConfig = new com.corundumstudio.socketio.Configuration();
        socketIOConfig.setHostname("0.0.0.0");
        socketIOConfig.setPort(7000);
        server = new SocketIOServer(socketIOConfig);

        server.start();

        server.addConnectListener(socketIOClient -> System.out.println("client connect: " + socketIOClient.getSessionId()));
        server.addDisconnectListener(socketIOClient -> System.out.println("client disconnect: " + socketIOClient.getSessionId()));

        server.addEventListener("order-status", String.class, ((socketIOClient, email, ackRequest) -> {
            server.getBroadcastOperations().sendEvent("order-status", email);
        }));

        return server;
    }

    @PreDestroy
    public void preDistroy() {
        server.stop();
    }
}
