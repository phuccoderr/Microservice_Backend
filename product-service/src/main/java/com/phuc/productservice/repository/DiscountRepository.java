package com.phuc.productservice.repository;

import com.phuc.productservice.models.Discount;
import com.phuc.productservice.models.Product;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.JpaSpecificationExecutor;
import org.springframework.stereotype.Repository;

@Repository
public interface DiscountRepository extends JpaRepository<Discount, String>, JpaSpecificationExecutor<Discount> {

    Discount findByCode(String code);
}
