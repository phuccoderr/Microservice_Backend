package com.phuc.categoryservice.repository;

import com.phuc.categoryservice.models.Category;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.UUID;

@Repository
public interface CategoryRepository extends JpaRepository<Category, String> {

    Category findByName(String name);
}
