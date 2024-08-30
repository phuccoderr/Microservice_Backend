package com.phuc.productservice.repository;

import com.phuc.productservice.models.Product;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.JpaSpecificationExecutor;
import org.springframework.data.jpa.repository.Modifying;
import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;

@Repository
public interface ProductRepository extends JpaRepository<Product, String>, JpaSpecificationExecutor<Product> {

    @Query("SELECT p FROM Product p WHERE p.name like %?1%")
    Page<Product> search(String keyword, Pageable pageable);
    Product findByName(String name);

    @Modifying
    @Query("UPDATE Product p SET p.reviewCount = p.reviewCount + 1, p.averageRating = ?2 WHERE p.id = ?1")
    void updateAverageRating(String proId, Float avgRating);


}
