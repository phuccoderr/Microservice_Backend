package com.phuc.productservice.initialize;

import com.corundumstudio.socketio.SocketIOServer;
import lombok.RequiredArgsConstructor;
import org.springframework.boot.CommandLineRunner;
import org.springframework.stereotype.Component;

@Component
@RequiredArgsConstructor
public class SocketIORunner implements CommandLineRunner {
    private final SocketIOServer server;

    @Override
    public void run(String... args) throws Exception {
        server.start();
        Runtime.getRuntime().addShutdownHook(new Thread(server::stop)); // Dừng server khi JVM tắt
    }
}
