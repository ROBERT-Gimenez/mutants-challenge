
# Instrucciones para Levantar el Proyecto

Para levantar este proyecto en tu máquina local, sigue los pasos a continuación:

## 1. Levantar Instala las dependencias (si tienes un go.mod)

```bash
go mod tidy
```

## 2. Crear archivo `.env`

Para configurar la conexión a DynamoDB, debes crear un archivo `.env` en la raíz del proyecto que contenga las siguientes variables de entorno:

### Variables de entorno para DynamoDB:

```env
AWS_ACCESS_KEY_ID=<tu_access_key>
AWS_SECRET_ACCESS_KEY=<tu_secret_key>
AWS_REGION=<tu_region>
```

## 3. Levantar y Ejecutar el Proyecto

go run main.go


# API Documentation

Este proyecto expone dos APIs principales para interactuar con datos de ADN y determinar si una secuencia corresponde a un mutante. A continuación se describen las dos APIs disponibles:


### 1. Registrar Mutante (POST)
- **Descripción:** Este endpoint permite detectar su el adn que ingresado es de un mutante o humano , en caso de no se un mutante retorna un 403 forbiden
- **Ruta:** `/mutant`
- **Método:** POST


#### Request Body:
```json

{
    "adn":["AAAADA" , "AWASES" , "ASDASA" , "ASSDAA", "QSWWDD", "ZSXXCD"]
}

```

### **Response (200 OK)**
```json
    true
```

### 1. Obtener Stats de Mutantes (GET)
- **Descripción:** Esta API permite obtener estadísticas relacionadas con el análisis de ADN. Retorna el conteo total de ADN de mutantes y humanos, junto con la relación entre ambos.
- **Ruta:** `/stats`
- **Método:** GET

### **Response (200 OK)**
```json
{
  "count_mutant_dna": 9,
  "count_human_dna": 19,
  "ratio": 0.47
}