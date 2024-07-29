package com.phuc.productservice.service;


import com.phuc.productservice.dtos.CategoryDto;
import com.phuc.productservice.dtos.CloudinaryDto;
import com.phuc.productservice.exceptions.DataErrorException;
import com.phuc.productservice.models.Product;
import com.phuc.productservice.repository.ProductRepository;
import com.phuc.productservice.request.RequestCreateProduct;
import com.phuc.productservice.util.Utility;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

@Service
@RequiredArgsConstructor
public class ProductService {
    
    private final ProductRepository repository;
    public Product createProduct(RequestCreateProduct reqCreateProduct, CategoryDto categoryDto) {

        Product product = Product.builder()
                .name(reqCreateProduct.getName())
                .alias(Utility.unAccent(reqCreateProduct.getName()))
                .description(reqCreateProduct.getDescription())
                .status(reqCreateProduct.getStatus())
                .cost(reqCreateProduct.getCost())
                .price(reqCreateProduct.getPrice())
                .sale(reqCreateProduct.getSale())
                .imageId(reqCreateProduct.getImageId())
                .url(reqCreateProduct.getUrl())
                .categoryId(categoryDto.getId())
                .build();

        if (reqCreateProduct.getUrl().isEmpty()) {
            product.setUrl("image default");
        }

        if (reqCreateProduct.getExtraImages().size() > 0) {
            for (CloudinaryDto cloud : reqCreateProduct.getExtraImages()) {
                product.addImage(cloud);
            }
        }

        return repository.save(product);
    }

    public void checkNameUnique(String name) throws DataErrorException {
        Product product = repository.findByName(name);
        if (product != null) {
            throw new DataErrorException("Data already exists!");
        }
    }
}
