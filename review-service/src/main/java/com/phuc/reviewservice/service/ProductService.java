package com.phuc.reviewservice.service;

import com.phuc.reviewservice.response.ResponseObject;
import jakarta.servlet.http.HttpServletRequest;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestTemplate;

@Service
@RequiredArgsConstructor
public class ProductService implements IProductService{
    private final RestTemplate restTemplate;
    private final String PRODUCT_SERVICE_URL = "http://localhost:9140/api/v1/products/";
    @Override
    public void findById(String proId, HttpServletRequest request) {
        String url = PRODUCT_SERVICE_URL + proId;

        restTemplate.getForEntity(url, ResponseObject.class);

    }
}
