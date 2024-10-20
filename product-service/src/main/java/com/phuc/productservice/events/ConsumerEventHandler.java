package com.phuc.productservice.events;


import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.phuc.productservice.events.message.AvgRatingEvent;
import com.phuc.productservice.events.message.PlaceOrderEvent;
import com.phuc.productservice.models.Product;
import com.phuc.productservice.repository.ProductRepository;
import lombok.RequiredArgsConstructor;
import org.apache.kafka.clients.consumer.ConsumerRecord;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.messaging.handler.annotation.Payload;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.io.IOException;
import java.util.Optional;


@Service
@RequiredArgsConstructor
public class ConsumerEventHandler {

    private final ProductRepository productRepository;
    private final ObjectMapper mapper;

    private final Logger LOGGER = LoggerFactory.getLogger(this.getClass());


    @KafkaListener(topics = "place-order-events-topics",
            groupId = "group-placeOrder",
            containerFactory = "placeOrderKafkaListenerContainerFactory")
    public void listener(PlaceOrderEvent placeOrderEvent) {
        LOGGER.info("Received message: {}", placeOrderEvent);

        placeOrderEvent.getOrderDetails().forEach( item -> {
            String proId = item.getProductId().getId();
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

    @KafkaListener(topics = "average-rating-events-topics",
            groupId = "group-avg")
    @Transactional
    public void listener(ConsumerRecord<String, byte[]> message) throws IOException {
        LOGGER.info("message is: {}",message.value());
        AvgRatingEvent avgRatingEvent = mapper.readValue(message.value(), AvgRatingEvent.class);
        LOGGER.info("ProId: {}, Average Rating: {}", avgRatingEvent.getProductId(),avgRatingEvent.getAverageRating());
        productRepository.updateAverageRating(avgRatingEvent.getProductId(),avgRatingEvent.getAverageRating());
    }
}
