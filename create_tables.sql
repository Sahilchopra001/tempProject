CREATE TABLE IF NOT EXISTS Library (
    ID INT PRIMARY KEY,
    Name VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS Users (
    ID INT PRIMARY KEY,
    Name VARCHAR(255),
    Email VARCHAR(255),
    ContactNumber VARCHAR(20),
    Role VARCHAR(20),
    LibID INT,
    FOREIGN KEY (LibID) REFERENCES Library(ID)
);

-- Define other tables similarly
