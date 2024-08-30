package com.phuc.reviewservice.configuration;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.client.RestTemplate;

@Configuration
public class WebApplicationConfig {
    @Bean
    public RestTemplate restTemplate() {
        return new RestTemplate();
    }
}
