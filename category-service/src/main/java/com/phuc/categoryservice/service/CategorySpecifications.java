package com.phuc.categoryservice.service;

import com.phuc.categoryservice.models.Category;
import jakarta.persistence.criteria.JoinType;
import jakarta.persistence.criteria.Predicate;
import org.springframework.data.jpa.domain.Specification;

public class CategorySpecifications {

    public static Specification<Category> defaultRootCategories() {
        return (root, query, criteriaBuilder) -> {
            query.distinct(true); // Đảm bảo không có các danh mục trùng lặp được trả về
            root.fetch("children", JoinType.LEFT); // Lấy các con một cách tự nhiên
            return criteriaBuilder.isNull(root.get("parent")); // Chỉ lấy các danh mục gốc
        };
    }

    public static Specification<Category> searchWithName(String keyword) {
        return (root, query, builder) -> {
            Predicate spec = builder.like(root.get("name"),"%" + keyword + "%");
            return builder.and(spec);
        } ;
    }
}
