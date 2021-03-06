#include <boost/test/unit_test.hpp>

#include <nav_parent_vs_value_inst.h>

#include <iostream>
#include <fstream>
#include <vector>

BOOST_AUTO_TEST_CASE(test_nav_parent_vs_value_inst) {
    std::ifstream ifs("src/term_strz.bin", std::ifstream::binary);
    kaitai::kstream ks(&ifs);
    nav_parent_vs_value_inst_t* r = new nav_parent_vs_value_inst_t(&ks);

    BOOST_CHECK_EQUAL(r->s1(), "foo");

    delete r;
}
