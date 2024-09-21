package com.phuc.productservice.service;

import com.corundumstudio.socketio.SocketIOServer;
import com.corundumstudio.socketio.annotation.OnConnect;
import com.corundumstudio.socketio.annotation.OnDisconnect;
import lombok.RequiredArgsConstructor;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Service;

@Service
@RequiredArgsConstructor
public class SocketIOService implements ISocketIOService{
    private static final Logger LOGGER = LoggerFactory.getLogger(SocketIOService.class);

    private final SocketIOServer server;

    @OnConnect
    public void onConnect() {
        LOGGER.info("Client connected");
    }

    @OnDisconnect
    public void onDisconnect() {
        LOGGER.info("Client disconnected");
    }

    @Override
    public void sendMessageToAddImage(String message) {
        server.getBroadcastOperations().sendEvent("add-image", message);
    }
}
