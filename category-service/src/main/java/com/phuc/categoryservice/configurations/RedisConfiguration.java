package com.phuc.categoryservice.configurations;

import com.phuc.categoryservice.models.Category;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.data.redis.connection.RedisConnectionFactory;
import org.springframework.data.redis.core.RedisTemplate;

@Configuration
public class RedisConfiguration {

    @Bean
    public RedisTemplate<Long, Category> redisTemplate(RedisConnectionFactory connectionFactory) {
        RedisTemplate<Long, Category> template = new RedisTemplate<>();
        template.setConnectionFactory(connectionFactory);
        return template;
    }
}
