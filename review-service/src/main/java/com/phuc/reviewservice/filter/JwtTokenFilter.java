package com.phuc.reviewservice.filter;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.phuc.reviewservice.constants.Constants;
import com.phuc.reviewservice.exeptions.InvalidTokenException;
import com.phuc.reviewservice.exeptions.TokenExpiredException;
import com.phuc.reviewservice.response.ResponseError;
import com.phuc.reviewservice.util.JwtTokenUtil;
import io.jsonwebtoken.Claims;
import jakarta.servlet.FilterChain;
import jakarta.servlet.ServletException;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import lombok.NonNull;
import lombok.RequiredArgsConstructor;
import org.modelmapper.internal.Pair;
import org.springframework.http.HttpStatus;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.authority.SimpleGrantedAuthority;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.web.authentication.WebAuthenticationDetailsSource;
import org.springframework.stereotype.Component;
import org.springframework.web.filter.OncePerRequestFilter;

import java.io.IOException;
import java.util.*;

@Component
@RequiredArgsConstructor
public class JwtTokenFilter extends OncePerRequestFilter {

    private final JwtTokenUtil jwtTokenUtil;

    private final ObjectMapper objectMapper = new ObjectMapper();

    @Override
    protected void doFilterInternal(HttpServletRequest request,
                                    HttpServletResponse response,
                                    FilterChain filterChain) throws ServletException, IOException {

        if (isBypassToken(request)) {
            filterChain.doFilter(request,response);
            return;
        }

        String authHeader = request.getHeader("Authorization");
        String email = null;
        String name = null;
        Set<String> roles = new HashSet<>();

        if (authHeader != null && authHeader.startsWith("Bearer ")) {
            String token = authHeader.substring(7);
            try {
                Claims claims = jwtTokenUtil.extractToken(token);
                email = claims.get("email").toString();
                name = claims.get("name").toString();
                List<String> rolesList = (List<String>) claims.get("roles");

                roles = new HashSet<>(rolesList);
            } catch (TokenExpiredException | InvalidTokenException ex) {
                handleException(response, ex.getMessage());
                return;
            }

        } else {
            handleException(response, Constants.TOKEN_INVALID);
        }

        Map<String, String> principal = new HashMap<>();
        principal.put("email",email);
        principal.put("name", name);

        if (email != null && SecurityContextHolder.getContext().getAuthentication() == null) {
            UsernamePasswordAuthenticationToken authToken  =
                    new UsernamePasswordAuthenticationToken(
                            principal,
                            null,
                            roles.stream().map(SimpleGrantedAuthority::new).toList());

            authToken.setDetails(new WebAuthenticationDetailsSource().buildDetails(request));
            SecurityContextHolder.getContext().setAuthentication(authToken);
        }

        filterChain.doFilter(request,response);
    }

    private void handleException(HttpServletResponse response, String message) throws IOException {
        response.setContentType("application/json");
        response.setStatus(HttpServletResponse.SC_UNAUTHORIZED);

        ResponseError errorResponse = ResponseError.builder()
                .message(List.of(message))
                .error(HttpStatus.UNAUTHORIZED.getReasonPhrase())
                .statusCode(HttpStatus.UNAUTHORIZED.value()).build();

        String jsonResponse = objectMapper.writeValueAsString(errorResponse);
        response.getWriter().write(jsonResponse);
        response.getWriter().flush();
    }

    private boolean isBypassToken(@NonNull HttpServletRequest request) {
        final List<Pair<String,String>> bypassToken = Arrays.asList(
                Pair.of(String.format("%s/parent/[^/]+",Constants.API_CATEGORIES), "GET")
        );

        String requestUri = request.getRequestURI();
        String requestMethod = request.getMethod();

        for (Pair<String, String> pair : bypassToken) {
            String uriPattern = pair.getLeft();
            String method = pair.getRight();

            if (requestUri.matches(uriPattern) && requestMethod.equalsIgnoreCase(method)) {
                return true;
            }
        }

        return false;
    }

}
