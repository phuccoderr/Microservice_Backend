package com.phuc.categoryservice.filter;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.phuc.categoryservice.response.ResponseError;
import com.phuc.categoryservice.util.JwtTokenUtil;
import io.jsonwebtoken.Claims;
import io.jsonwebtoken.ExpiredJwtException;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.authority.SimpleGrantedAuthority;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.web.authentication.WebAuthenticationDetailsSource;
import org.springframework.stereotype.Component;
import org.springframework.web.ErrorResponse;
import org.springframework.web.filter.OncePerRequestFilter;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import jakarta.servlet.FilterChain;
import jakarta.servlet.ServletException;

import java.io.IOException;
import java.util.*;
import java.util.stream.Collectors;


@Component
public class JwtTokenFilter extends OncePerRequestFilter {

    @Autowired
    JwtTokenUtil jwtTokenUtil;

    private final ObjectMapper objectMapper = new ObjectMapper();

    @Override
    protected void doFilterInternal(HttpServletRequest request,
                                    HttpServletResponse response,
                                    FilterChain filterChain) throws ServletException, IOException {
        try {
            String authHeader = request.getHeader("Authorization");
            String token = null;
            String email = null;
            Set<String> roles = new HashSet<>();

            if (authHeader != null && authHeader.startsWith("Bearer ")) {
                token = authHeader.substring(7);
                Claims claims = jwtTokenUtil.extractToken(token);
                email = claims.get("email").toString();
                List<String> rolesList = (List<String>) claims.get("roles");

                roles = new HashSet<>(rolesList);
            }


            if (email != null && SecurityContextHolder.getContext().getAuthentication() == null) {
                UsernamePasswordAuthenticationToken authToken  =
                        new UsernamePasswordAuthenticationToken(
                                email,
                                null,
                                roles.stream().map(SimpleGrantedAuthority::new).collect(Collectors.toList()));

                authToken.setDetails(new WebAuthenticationDetailsSource().buildDetails(request));
                SecurityContextHolder.getContext().setAuthentication(authToken);
            }
        } catch (ExpiredJwtException e) {
            response.setContentType("application/json");
            response.setStatus(HttpServletResponse.SC_UNAUTHORIZED);

            ResponseError errorResponse = ResponseError.builder()
                    .message(List.of("Authentication JWT Expired Or Signature"))
                    .error(HttpStatus.UNAUTHORIZED.getReasonPhrase())
                    .statusCode(HttpStatus.UNAUTHORIZED.value()).build();

            String jsonResponse = objectMapper.writeValueAsString(errorResponse);
            response.getWriter().write(jsonResponse);
            response.getWriter().flush();
            return;
        }

        filterChain.doFilter(request,response);
    }

}
