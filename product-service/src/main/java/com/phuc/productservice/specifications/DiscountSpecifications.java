package com.phuc.productservice.specifications;

import com.phuc.productservice.models.Discount;
import jakarta.persistence.criteria.Predicate;
import org.springframework.data.jpa.domain.Specification;

public class DiscountSpecifications {

    public static Specification<Discount> withKeyword(String keyword) {
        return (root, query, criteriaBuilder) -> {
            Predicate search = criteriaBuilder.or(
                    criteriaBuilder.like(root.get("name"), "%" + keyword + "%"),
                    criteriaBuilder.like(root.get("code"), "%" + keyword + "%")
            );
            return criteriaBuilder.and(search);
        };
    }
}
