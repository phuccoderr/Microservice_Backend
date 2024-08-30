package com.phuc.reviewservice.repository;

import com.phuc.reviewservice.models.Review;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;

@Repository
public interface ReviewRepository extends JpaRepository<Review, String> {

    @Query("SELECT AVG(r.rating) FROM Review r WHERE r.productId = ?1")
    Float findAvgRating(String proId);

    @Query("SELECT r FROM Review r WHERE r.headline LIKE %?1% OR r.comment LIKE %?1%")
    Page<Review> search(String keyword, Pageable pageable);

    @Query("SELECT r FROM Review r WHERE r.productId = ?1")
    Page<Review> findByProduct(String proId, Pageable pageable);

}
