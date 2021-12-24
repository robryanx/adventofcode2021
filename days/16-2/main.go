package main

import (
    "fmt"
    "strings"
    "strconv"
    "adventofcode/2021/modules/readinput"
)

type packet_s struct {
    bits string
    pos int
    state string
    version int64
    type_id int64
    numbers string
    number_value int64

    sub_packets []*packet_s
}

func main() {
    bits_lookup := map[rune]string{
        '0': "0000",
        '1': "0001",
        '2': "0010",
        '3': "0011",
        '4': "0100",
        '5': "0101",
        '6': "0110",
        '7': "0111",
        '8': "1000",
        '9': "1001",
        'A': "1010",
        'B': "1011",
        'C': "1100",
        'D': "1101",
        'E': "1110",
        'F': "1111",  
    }

    input := readinput.ReadStrings("inputs/16/input.txt", "\n")[0]

    bits := ""
    for _, hex_val := range strings.Split(input, "") {
        bits += bits_lookup[rune(hex_val[0])]
    }

    parent_packet := &packet_s {
        state: "parent",
    }

    packet := &packet_s{
        bits: bits,
        state: "version",
    }

    parent_packet, _ = parse_packet(parent_packet, packet, -1)

    fmt.Println(operate_packet(parent_packet, 1))
}

func draw_tree(packet *packet_s, level int) {
    for i:=0; i<level*4; i++ {
        fmt.Printf("=")
    }

    fmt.Printf("version=%d, state=%s, type_id=%d, numbers=%s, number_value=%d\n", packet.version, packet.state, packet.type_id, packet.numbers, packet.number_value)

    for i:=0; i<len(packet.sub_packets); i++ {
        draw_tree(packet.sub_packets[i], (level+1))
    }
}

func operate_packet(packet *packet_s, level int) int64 {
    if len(packet.sub_packets) == 0 {
        return packet.number_value
    }

    value := int64(0)

    switch packet.type_id {
    case 0:
        for i:=0; i<len(packet.sub_packets); i++ {
            value += operate_packet(packet.sub_packets[i], (level+1))
        }
    case 1:
        value = 1
        for i:=0; i<len(packet.sub_packets); i++ {
            value *= operate_packet(packet.sub_packets[i], (level+1))
        }
    case 2:
        value = int64(10000000000000)
        for i:=0; i<len(packet.sub_packets); i++ {
            sub_value := operate_packet(packet.sub_packets[i], (level+1))
            if sub_value < value {
                value = sub_value
            }
        }
    case 3:
        value = int64(0)
        for i:=0; i<len(packet.sub_packets); i++ {
            sub_value := operate_packet(packet.sub_packets[i], (level+1))
            if sub_value > value {
                value = sub_value
            }
        }
    case 5:
        first_value := operate_packet(packet.sub_packets[0], (level+1))
        second_value := operate_packet(packet.sub_packets[1], (level+1))

        if first_value > second_value {
            value = 1
        }
    case 6:
        first_value := operate_packet(packet.sub_packets[0], (level+1))
        second_value := operate_packet(packet.sub_packets[1], (level+1))

        if first_value < second_value {
            value = 1
        }
    case 7:
        first_value := operate_packet(packet.sub_packets[0], (level+1))
        second_value := operate_packet(packet.sub_packets[1], (level+1))

        if first_value == second_value {
            value = 1
        }
    }

    return value
}

func parse_packet(parent_packet *packet_s, packet *packet_s, count int) (*packet_s, int) {
    for ;; {
        if packet.state == "version" && packet.pos + 3 < len(packet.bits) {
            packet.version, _ = strconv.ParseInt(packet.bits[packet.pos:packet.pos+3], 2, 64)
            packet.state = "type"
            packet.pos += 3
        } else if packet.state == "type" && packet.pos + 3 < len(packet.bits) {
            packet.type_id, _ = strconv.ParseInt(packet.bits[packet.pos:packet.pos+3], 2, 64)
            packet.pos += 3

            if packet.type_id == 4 {
                packet.state = "numbers"
            } else {
                if packet.bits[packet.pos] == '0' {
                    packet.state = "sub_packet_length"
                } else {
                    packet.state = "sub_packet_count"
                }

                packet.pos += 1
            }
        } else if packet.state == "numbers" {
            packet.numbers += packet.bits[packet.pos+1:packet.pos+5]

            if packet.bits[packet.pos] == '0' {
                packet.number_value, _ = strconv.ParseInt(packet.numbers, 2, 64)

                packet.pos += 5

                break;
            }

            packet.pos += 5
        } else if packet.state == "sub_packet_length" {
            length, _ := strconv.ParseInt(packet.bits[packet.pos:packet.pos+15], 2, 64)

            packet.pos += 15

            next_packet := &packet_s{
                bits: packet.bits[packet.pos:packet.pos+int(length)],
                state: "version",
            }

            packet.pos += int(length)
            
            parse_packet(packet, next_packet, -1)

            break;
        } else if packet.state == "sub_packet_count" {
            count, _ := strconv.ParseInt(packet.bits[packet.pos:packet.pos+11], 2, 64)
            packet.pos += 11

            next_packet := &packet_s{
                bits: packet.bits[packet.pos:],
                state: "version",
            }
            
            _, pos := parse_packet(packet, next_packet, int(count))

            packet.pos += pos

            break;
        }
    }

    parent_packet.sub_packets = append(parent_packet.sub_packets, packet)
    if count > -1 && len(parent_packet.sub_packets) == count {
        return parent_packet, packet.pos
    }

    // are there more packets?
    if len(packet.bits)-packet.pos > 10 {
        next_packet := &packet_s{
            bits: packet.bits[packet.pos:],
            state: "version",
        }

        _, pos := parse_packet(parent_packet, next_packet, count)
        packet.pos += pos

        if count > -1 && len(parent_packet.sub_packets) == count {
            return parent_packet, packet.pos
        }
    }

    return parent_packet, packet.pos
}