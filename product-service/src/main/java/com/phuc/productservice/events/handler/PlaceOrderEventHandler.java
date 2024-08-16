package com.phuc.productservice.events.handler;


import com.phuc.productservice.events.message.PlaceOrderEvent;
import com.phuc.productservice.models.Product;
import com.phuc.productservice.repository.ProductRepository;
import lombok.RequiredArgsConstructor;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Service;

import java.util.Optional;


@Service
@RequiredArgsConstructor
public class PlaceOrderEventHandler {

    private final ProductRepository productRepository;

    private final Logger LOGGER = LoggerFactory.getLogger(this.getClass());


    @KafkaListener(topics = "place-order-events-topics", groupId = "groupA", containerFactory = "placeOrderKafkaListenerContainerFactory")
    public void listen(PlaceOrderEvent placeOrderEvent) {

        placeOrderEvent.getItems().forEach( item -> {
            String proId = item.getProductId();
            int quantity = item.getQuantity();

            Optional<Product> productOptional  = productRepository.findById(proId);
            if (productOptional.isPresent()) {
                Product product = productOptional.get();

                if (product.getStock() >= quantity) {
                    product.setStock(product.getStock() - quantity);
                    productRepository.save(product);

                    LOGGER.info("Stock deducted for product ID: {}, new stock: {}", proId, product.getStock());
                } else {
                    LOGGER.warn("Not enough stock for product ID: {}, requested: {}, available: {}",
                            proId, quantity, product.getStock());
                }
            } else {
                LOGGER.error("Product ID: {} not found in the repository", proId);
            }
        });
    }
}
