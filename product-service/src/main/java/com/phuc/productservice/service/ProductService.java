package com.phuc.productservice.service;


import com.phuc.productservice.dtos.CategoryDto;
import com.phuc.productservice.dtos.CloudinaryDto;
import com.phuc.productservice.exceptions.DataErrorException;
import com.phuc.productservice.exceptions.ParamValidateException;
import com.phuc.productservice.models.Product;
import com.phuc.productservice.repository.ProductRepository;
import com.phuc.productservice.request.RequestCreateProduct;
import com.phuc.productservice.request.RequestProduct;
import com.phuc.productservice.request.RequestUpdateProduct;
import com.phuc.productservice.util.Utility;
import lombok.RequiredArgsConstructor;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;
import org.springframework.data.domain.Pageable;
import org.springframework.data.domain.Sort;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.zip.DataFormatException;

@Service
@RequiredArgsConstructor
public class ProductService implements IProductService {
    
    private final ProductRepository repository;

    @Override
    public Page<Product> getAllProducts(Integer page, Integer limit, String sort, String keyword)
            throws ParamValidateException {

        Utility.checkSortIsAscOrDesc(sort);

        Sort sortDir = Sort.by("name");
        sortDir = sort.equals("asc") ? sortDir.ascending() : sortDir.descending();

        Pageable pageable = PageRequest.of(page - 1,limit,sortDir);

        if (!keyword.isEmpty()) {
            return repository.search(keyword, pageable);
        } else  {
            return repository.findAll(pageable);
        }
    }

    @Override
    public Product getProduct(String proId) throws DataErrorException {
        return repository.findById(proId).orElseThrow( () -> new DataErrorException("Data not Found!") );
    }

    @Override
    public Product createProduct(Product product,RequestProduct requestProduct, CategoryDto categoryDto) {

        setDtoToEntity(product,requestProduct);

        if(categoryDto != null) {
            product.setCategoryId(categoryDto.getId());
        }

        return repository.save(product);
    }

    public Product updateProduct(Product proInDB,RequestProduct requestProduct, CategoryDto categoryDto) throws DataErrorException {

        setDtoToEntity(proInDB,requestProduct);

        if(categoryDto != null) {
            proInDB.setCategoryId(categoryDto.getId());
        }

        return repository.save(proInDB);
    }

    @Override
    public void checkNameUnique(String name) throws DataErrorException {
        Product product = repository.findByName(name);
        if (product != null) {
            throw new DataErrorException("Data already exists!");
        }
    }

    public void checkNameUnique(String oldName, String newName) throws DataErrorException {
        if (!oldName.equals(newName)) {
            checkNameUnique(newName);
        }
    }

    private void setDtoToEntity(Product product, RequestProduct reqProduct) {
        product.setName(reqProduct.getName());
        product.setAlias(Utility.unAccent(reqProduct.getName()));
        product.setDescription(reqProduct.getDescription());
        product.setStatus(reqProduct.getStatus());
        product.setCost(reqProduct.getCost());
        product.setPrice(reqProduct.getPrice());
        product.setSale(reqProduct.getSale());
    }
}
