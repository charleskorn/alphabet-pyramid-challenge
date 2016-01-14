# alphabet-pyramid-challenge

For all your text-based pyramid needs.

    $ ./go.sh run K
              A
             B B
            C   C
           D     D
          E       E
         F         F
        G           G
       H             H
      I               I
     J                 J
    K                   K
     J                 J
      I               I
       H             H
        G           G
         F         F
          E       E
           D     D
            C   C
             B B
              A

## Prerequisites

* Go 1.5.3

Run `./go.sh check` to test that everything looks OK.

## Building

* Run `./go.sh build` to build

## Running

* Build application
* Run `./go.sh run <letter>` to run the application (replace `<letter>` with your letter of choice)

## Testing

* Run `./go.sh setup` to pull down dependencies if you haven't already (only need to do this once)
* Run `./go.sh test` to build and run the tests
