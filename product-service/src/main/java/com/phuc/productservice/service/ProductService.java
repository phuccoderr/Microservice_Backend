package com.phuc.productservice.service;

import com.phuc.productservice.constants.Constants;
import com.phuc.productservice.dtos.CategoryDto;
import com.phuc.productservice.dtos.CloudinaryDto;
import com.phuc.productservice.exceptions.*;
import com.phuc.productservice.models.Product;
import com.phuc.productservice.models.ProductImage;
import com.phuc.productservice.repository.ProductImageRepository;
import com.phuc.productservice.repository.ProductRepository;
import com.phuc.productservice.request.RequestProduct;
import com.phuc.productservice.util.FileUploadUtil;
import com.phuc.productservice.util.Utility;
import lombok.RequiredArgsConstructor;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;
import org.springframework.data.domain.Pageable;
import org.springframework.data.domain.Sort;
import org.springframework.data.jpa.domain.Specification;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.web.multipart.MultipartFile;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.concurrent.CompletableFuture;

@Service
@RequiredArgsConstructor
public class ProductService implements IProductService {

    private static final Logger LOGGER = LoggerFactory.getLogger(ProductService.class);
    
    private final ProductRepository proRepository;
    private final ProductImageRepository proImageRepository;
    private final CloudinaryService cloudinaryService;
    private final SocketIOService socketIOService;
    private final ProductRedisService productRedisService;

    @Override
    public Page<Product> getAllProducts(Integer page, Integer limit, String sort, String keyword)
            throws ParamValidateException {

        Utility.checkSortIsAscOrDesc(sort);

        Sort sortDir = Sort.by("name");
        sortDir = sort.equals("asc") ? sortDir.ascending() : sortDir.descending();

        Pageable pageable = PageRequest.of(page - 1,limit,sortDir);

        if (!keyword.isEmpty()) {
            return proRepository.search(keyword, pageable);
        } else  {
            return proRepository.findAll(pageable);
        }
    }

    @Override
    public Product getProduct(String proId) throws DataNotFoundException {
        return proRepository.findById(proId).orElseThrow( () -> new DataNotFoundException(Constants.DB_NOT_FOUND) );
    }



    @Override
    public Product createProduct(
            MultipartFile mainImage,
            List<MultipartFile> extraImages,
            RequestProduct requestProduct,
            CategoryDto categoryDto
    ) throws FuncErrorException {

        Product product = new Product();
        setDtoToEntity(product,requestProduct);

        if (mainImage != null && !mainImage.isEmpty()) {
            setMainImage(mainImage,product);
        }

        if(categoryDto != null) {
            product.setCategoryId(categoryDto.getId());
        }

        return proRepository.save(product);
    }

    @Override
    public Product updateProduct(
            Product proInDB,
            RequestProduct requestProduct,
            CategoryDto categoryDto,
            MultipartFile mainImage
    ) throws FuncErrorException {

        setDtoToEntity(proInDB,requestProduct);

        if (mainImage != null && !mainImage.isEmpty()) {
            cloudinaryService.deleteAsyncImage(proInDB.getImageId());
            setMainImage(mainImage,proInDB);
        }

        if(categoryDto != null) {
            proInDB.setCategoryId(categoryDto.getId());
        }

        return proRepository.save(proInDB);
    }

    @Override
    public void deleteProductById(String proId) throws DataNotFoundException {
        Product product = getProduct(proId);
        proRepository.delete(product);

        cloudinaryService.deleteAsyncImage(product.getImageId());
        product.getExtraImages().forEach( item -> {
            cloudinaryService.deleteAsyncImage(item.getImageId());
        });
    }

    @Override
    public void checkNameUnique(String name) throws DataErrorException {
        Product product = proRepository.findByName(name);
        if (product != null) {
            throw new DataErrorException(Constants.DB_ALREADY_EXISTS);
        }
    }

    @Override
    public void checkNameUnique(String oldName, String newName) throws DataErrorException {
        if (!oldName.equals(newName)) {
            checkNameUnique(newName);
        }
    }

    @Override
    public Product addFiles(List<MultipartFile> extraFiles, Product product) throws FuncErrorException {
        setExtraImage(extraFiles,product);
        for (MultipartFile file : extraFiles) {
            FileUploadUtil.assertAllowed(file, FileUploadUtil.IMAGE_PATTERN);

            String fileName = FileUploadUtil.getFileName(file.getOriginalFilename());

            CloudinaryDto cloudinaryDto = cloudinaryService.uploadImage(file, fileName);

            product.addImage(cloudinaryDto);
        }
        return proRepository.save(product);
    }

    @Override
    @Transactional
    public void deleteFiles(List<String> listFiles, Product product) {
        listFiles.forEach(fileId -> {
            ProductImage productImage = proImageRepository.findById(fileId).orElseThrow(
                    () -> new DataNotFoundException(Constants.DB_NOT_FOUND)
            );
            cloudinaryService.deleteAsyncImage(productImage.getImageId());
            proImageRepository.deleteByIdByProductId(productImage.getId(),product.getId());
        });
    }

    @Override
    public Page<Product> getAllProductsByCategory(
            List<String> listCategoryIds,
            Integer page,
            Integer limit,
            String sort,
            String sortField,
            String keyword
    ) throws ParamValidateException {
        Utility.checkSortIsAscOrDesc(sort);

        Specification<Product> spec = Specification.where(null);
        spec = spec.and(ProductSpecifications.withCategory(listCategoryIds));

        if (!keyword.isEmpty()) {
            spec = spec.and(ProductSpecifications.withKeyword(keyword));
        }

        Map<String, String> sortFields = new HashMap<>();
        sortFields.put("price", "price");
        sortFields.put("date", "createdAt");

        String sortFieldDB = sortFields.getOrDefault(sortField, sortField);

        Pageable pageable = PageRequest.of(page - 1,limit, Sort.by(
                sort.equals("asc") ? Sort.Direction.ASC : Sort.Direction.DESC,
                sortFieldDB
        ));

        return proRepository.findAll(spec,pageable);

    }

    @Override
    public void setExtraImage(List<MultipartFile> extraFile,Product product) {

        if (extraFile != null && !extraFile.isEmpty() ) {
            List<CompletableFuture<Void>> uploadFutures = new ArrayList<>();
            for (MultipartFile file : extraFile) {
                String fileName = FileUploadUtil.getFileName(file.getOriginalFilename());
                CompletableFuture<Void> future  = cloudinaryService.uploadImageAsync(file, fileName, product);
                uploadFutures.add(future);
            }

            CompletableFuture.allOf(uploadFutures.toArray(new CompletableFuture[0])).thenRun(() -> {
                socketIOService.sendMessageToAddImage(String.format("File của sản phẩm %s đã upload hoàn tất", product.getName()));
                productRedisService.clear();
            });
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
        product.setStock(reqProduct.getStock());
        product.setAverageRating(product.getAverageRating() != null ? product.getAverageRating() : 0f);
        product.setReviewCount(product.getReviewCount() != null ? product.getReviewCount() : 0);
    }

    private void setMainImage(MultipartFile mainFile,Product product)
            throws FuncErrorException {

        FileUploadUtil.assertAllowed(mainFile, FileUploadUtil.IMAGE_PATTERN);

        String fileName = FileUploadUtil.getFileName(mainFile.getOriginalFilename());

        CloudinaryDto result = cloudinaryService.uploadImage(mainFile, fileName);

        product.setImageId(result.getPublicId());
        product.setUrl(result.getUrl());

    }

}
