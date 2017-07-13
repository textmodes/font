# ROMs

Here you will find raw (EEP)ROM dumps for original fonts. Also sometimes referred
to as "raw" fonts. These files usually have the column stored in the bits, one
byte per row.

To allow larger fonts with widths up to 16 pixels, we also support big endian
encoded unsigned 16-bit values.

## Maps

To help import the ROMs properly, we use maps with the following opcodes:

    #         comment
    @wxh      glyph width (w) and height (h) integer values
    >1        right advance (added to width, can be negative) (number)
    +n        skip count, number of bytes to be skipped in beginning of ROM (number)
    -n        trim count, number of bytes to be skipped at the end of ROM (number)
    =0x20     replacement character (number)
    *0x20     offset, marks the first character (number)
    u0x0020   unicode code point (array of numbers)

