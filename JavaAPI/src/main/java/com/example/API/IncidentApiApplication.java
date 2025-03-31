package com.example.API;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import jakarta.persistence.*;
import java.time.LocalDateTime;
import java.util.List;
import java.util.Optional;

@Entity
//Clase provisional para una base de datos local, por el momento el status siempre es pendiente si no se modifica.  
class Incident {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    private String reporter;
    private String description;
    private String status = "pendiente";
    private LocalDateTime createdAt = LocalDateTime.now();

    public Long getId() {
        return id;
    }
    
    public void setId(Long id) {
        this.id = id;
    }
    
    public String getReporter() {
        return reporter;
    }
    
    public void setReporter(String reporter) {
        this.reporter = reporter;
    }
    
    public String getDescription() {
        return description;
    }
    
    public void setDescription(String description) {
        this.description = description;
    }
    
    public String getStatus() {
        return status;
    }
    
    public void setStatus(String status) {
        this.status = status;
    }
    
    public LocalDateTime getCreatedAt() {
        return createdAt;
    }
    
    public void setCreatedAt(LocalDateTime createdAt) {
        this.createdAt = createdAt;
    }
    
}

interface IncidentRepository extends JpaRepository<Incident, Long> {}

@RestController
//Mando como request principal el incidents para trabajar sobre ella.
@RequestMapping("/incidents")
class IncidentController {
    @Autowired
    private IncidentRepository repository;

    @PostMapping
    public ResponseEntity<?> createIncident(@RequestBody Incident incident) {
        //Sin reporter no se debe poder crear un incidente. 
        if (incident.getReporter() == null || incident.getReporter().isEmpty()) {
            return ResponseEntity.badRequest().body("El reporter es obligatorio");
        }
        //Condición de Ludwing 
        if (incident.getDescription() == null || incident.getDescription().length() < 10) {
            return ResponseEntity.badRequest().body("La descripción debe tener al menos 10 caracteres");
        }
        return ResponseEntity.ok(repository.save(incident));
    }

    @GetMapping
    public List<Incident> getAllIncidents() {
        return repository.findAll();
    }

    @GetMapping("/{id}")
    public ResponseEntity<?> getIncident(@PathVariable Long id) {
        Optional<Incident> incident = repository.findById(id);
        return incident.map(ResponseEntity::ok).orElseGet(() -> ResponseEntity.notFound().build());
    }

    @PutMapping("/{id}")
    public ResponseEntity<?> updateIncident(@PathVariable Long id, @RequestBody Incident updatedIncident) {
        return repository.findById(id).map(incident -> {
            incident.setStatus(updatedIncident.getStatus());
            repository.save(incident);
            return ResponseEntity.ok("Estado actualizado");
        }).orElseGet(() -> ResponseEntity.notFound().build());
    }

    @DeleteMapping("/{id}")
    public ResponseEntity<?> deleteIncident(@PathVariable Long id) {
        return repository.findById(id).map(incident -> {
            repository.delete(incident);
            return ResponseEntity.ok("Incidente eliminado");
        }).orElseGet(() -> ResponseEntity.notFound().build());
    }
}

@SpringBootApplication
public class IncidentApiApplication {
    public static void main(String[] args) {
        //Estaba usando el puerto 8080 para base de datos, lo pase al 8083
        System.setProperty("server.port", "8083");
        SpringApplication.run(IncidentApiApplication.class, args);
    }
}
