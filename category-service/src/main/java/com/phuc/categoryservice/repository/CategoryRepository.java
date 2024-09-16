package com.phuc.categoryservice.repository;

import com.phuc.categoryservice.models.Category;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.JpaSpecificationExecutor;
import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface CategoryRepository extends JpaRepository<Category, String>, JpaSpecificationExecutor<Category> {

    @Query("SELECT c FROM Category c WHERE c.name like %?1%")
    Page<Category> search(String keyword, Pageable pageable);

    @Query("SELECT c FROM Category c WHERE c.parent.id is null")
    public List<Category> findRootCategories();

    Category findByName(String name);
}
