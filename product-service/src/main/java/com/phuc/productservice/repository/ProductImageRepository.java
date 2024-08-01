package com.phuc.productservice.repository;

import com.phuc.productservice.models.ProductImage;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Modifying;
import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;
import org.springframework.transaction.annotation.Transactional;

@Repository
public interface ProductImageRepository extends JpaRepository<ProductImage, String> {

    @Transactional
    @Modifying
    @Query("DELETE ProductImage img WHERE img.product.id = ?2 AND img.id = ?1")
    void deleteByIdByProductId(String id, String proId);
}
