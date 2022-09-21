#[allow(clippy::all)]
mod integers_export {
  #[export_name = "res"]
  unsafe extern "C" fn __wit_bindgen_integers_export_res() -> i32{
    let result = <super::IntegersExport as IntegersExport>::res();
    wit_bindgen_rust::rt::as_i32(result)
  }
  pub trait IntegersExport {
    fn res() -> u8;
  }
}

// struct IntegersExport {}

// impl integers_export::IntegersExport for IntegersExport {
//   fn res() -> u8 {
//     42
//   }
// }

// IntegersExport::res()