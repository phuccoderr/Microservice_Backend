package com.phuc.productservice.service;

import com.phuc.productservice.models.Product;
import jakarta.persistence.criteria.CriteriaBuilder;
import jakarta.persistence.criteria.Predicate;
import org.springframework.data.jpa.domain.Specification;

import java.util.List;

public class ProductSpecifications {

    public static Specification<Product> withKeyword(String keyword) {
        return (root, query, criteriaBuilder) -> {
            Predicate enabled = criteriaBuilder.isTrue(root.get("status"));
            Predicate search = criteriaBuilder.or(
                    criteriaBuilder.like(root.get("name"), "%" + keyword + "%"),
                    criteriaBuilder.like(root.get("description"), "%" + keyword + "%")
            );
            return criteriaBuilder.and(enabled,search);
        };
    }

    public static Specification<Product> withCategory(List<String> categoryIds) {
        return (root, query, criteriaBuilder) ->  {
            Predicate enabled = criteriaBuilder.isTrue(root.get("status"));

            CriteriaBuilder.In<String> categories = criteriaBuilder.in(root.get("categoryId"));

            categoryIds.forEach(categories::value);

            return criteriaBuilder.and(enabled,categories);
        };
    }
}
