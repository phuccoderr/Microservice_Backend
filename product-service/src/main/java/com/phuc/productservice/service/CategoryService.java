package com.phuc.productservice.service;

import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.phuc.productservice.dtos.CategoryDto;
import com.phuc.productservice.response.ResponseObject;
import jakarta.servlet.http.HttpServletRequest;
import lombok.RequiredArgsConstructor;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpMethod;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestTemplate;

import java.util.List;

@Service
@RequiredArgsConstructor
public class CategoryService implements ICategoryService {

    private final ObjectMapper objectMapper;
    private final RestTemplate restTemplate;
    private final String CATEGORY_SERVICE_URL = "http://category-service:9130/api/v1/categories/";

    @Override
    public CategoryDto getCategoryById(String categoryId, HttpServletRequest request) {

        String url = CATEGORY_SERVICE_URL + categoryId;

        String jwtToken = request.getHeader("Authorization");

        HttpHeaders headers = new HttpHeaders();
        headers.set("Authorization",jwtToken);

        HttpEntity<String> entity = new HttpEntity<>(headers);

        ResponseEntity<ResponseObject> response = restTemplate.exchange(
                url,
                HttpMethod.GET,
                entity,
                ResponseObject.class
        );

        Object data = response.getBody().getData();
        return objectMapper.convertValue(data, CategoryDto.class);
    }

    @Override
    public List<String> getChildrenCateId(String id, HttpServletRequest request) {
        String url = CATEGORY_SERVICE_URL + "parent/" + id;

        HttpEntity<Void> entity = new HttpEntity<>(null);

        ResponseEntity<ResponseObject> response = restTemplate.exchange(
                url,
                HttpMethod.GET,
                entity,
                ResponseObject.class
        );

        Object data = response.getBody().getData();
        return objectMapper.convertValue(data, new TypeReference<List<String>>() {
        });
    }
}
