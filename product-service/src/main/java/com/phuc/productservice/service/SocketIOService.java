package com.phuc.productservice.service;

import com.corundumstudio.socketio.SocketIOClient;
import com.corundumstudio.socketio.SocketIOServer;
import lombok.RequiredArgsConstructor;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Service;

import java.util.UUID;

@Service
@RequiredArgsConstructor
public class SocketIOService implements ISocketIOService{
    private static final Logger LOGGER = LoggerFactory.getLogger(SocketIOService.class);

    private final SocketIOServer server;

    @Override
    public void sendMessageToAddImage(String sessionId,String message) {
        SocketIOClient client = server.getClient(UUID.fromString(sessionId));
        if (client != null) {
            client.sendEvent("add-image", message); // Gửi event "addImage" đến client cụ thể
        }
    }


}
