package com.phuc.productservice.configuations;

import com.cloudinary.Cloudinary;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import java.util.HashMap;
import java.util.Map;

@Configuration
public class CloudinaryConfiguration {

    @Value("${CLOUD_NAME}")
    private String cloudName;

    @Value("${CLOUD_KEY}")
    private String cloudKey;

    @Value("${CLOUD_SECRET}")
    private String cloudSecret;

    @Bean
    public Cloudinary getCloudinary() {
        Map<String,String> config = new HashMap<>();
        config.put("cloud_name",cloudName);
        config.put("api_key", cloudKey);
        config.put("api_secret", cloudSecret);
        config.put("secure", "true"); // Sử dụng HTTPS
        return new Cloudinary(config);
    }
}
