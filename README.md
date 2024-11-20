# Memory-store

Simple memory store implementation

## Usage 
### Compile

1. Clone this repostiory
2. Run with
   ```sh
   $ go run main.go
   ```

### Command
1. **Get**: 
   Comma and space-separated attributes. Example:

   Print "No entry found for <key> " if get returns null.
   ```sh
   get
   ```
   result
   ```bash
   attribute1: attribute_value_1, attribute2: attribute_value_2, attribute3: attribute_value_3
   ```
2. **Put**:  Do not print anything.
   ```sh
    put sde_bootcamp title SDE-Bootcamp price 30000.00 enrolled false estimated_time 30
    ```
3. **Delete**:  Do not print anything.
    ```sh
    delete sde_bootcamp
    ```
4. **Search**:  Comma-separated keys. Example:
    ```sh
    search price 30000.00
    ```
    result
    ```sh
    sde_bootcamp
    ```
5. **Keys**:  Comma-separated keys. Example:
    ```sh
    keys
    ````
    result
    ```sh
    sde_bootcamp,sde_kickstart
    ```



